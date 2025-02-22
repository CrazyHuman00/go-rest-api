import axios from "axios"
import { useQueryClient, useMutation } from "@tanstack/react-query"
import { Task } from "../types"
import useStore from "../store"
import { useError } from "../hooks/useError"

export const useMutateTask = () => {
    const queryClient = useQueryClient()
    const resetEditedTask = useStore((state) => state.resetEditedTask)
    const { switchErrorHandling } = useError()

    // タスクの追加処理
    const createTaskMutation = useMutation(
        {
            mutationFn: async (task: Omit<Task, 'id' | 'created_at' | 'updated_at' >) => {
                const res = axios.post<Task>(`${process.env.REACT_APP_API_URL}/tasks`, task)
                return res
            },
            onSuccess: (res) => {
                // キャッシュのデータを取得し、tasksのキーワードで取得したデータに新しいデータを追加
                const previousTasks = queryClient.getQueryData<Task[]>(['tasks'])
                if (previousTasks) {
                    // 配列の末尾に新しいデータを追加し、新しい配列に更新する
                    queryClient.setQueryData<Task[]>(['tasks'], [...previousTasks, res.data])
                }
                resetEditedTask()
            },
            onError: (error: any) => {
                if (error.response.data.message) {
                    switchErrorHandling(error.response.data.message)
                } else {
                    switchErrorHandling(error.response.data)
                }
            }
        }
    )

    // タスクの更新処理
    const updateTaskMutation = useMutation(
        {
            mutationFn: async (task: Omit<Task, 'created_at' | 'updated_at' >) => {
                const res = axios.put<Task>(`
                    ${process.env.REACT_APP_API_URL}/tasks/${task.id}`,
                    {title: task.title}
                )
                return res
            },
            onSuccess: (res, variables) => {
                // 既存のタスクを取得
                const previousTasks = queryClient.getQueryData<Task[]>(['tasks'])
                if (previousTasks) {
                    // 更新後のタスクに書き換える
                    queryClient.setQueryData<Task[]>(
                        ['tasks'],
                        previousTasks.map((task) => (task.id === variables.id ? res.data : task))
                    )
                }
                resetEditedTask()
            },
            onError: (error: any) => {
                if (error.response.data.message) {
                    switchErrorHandling(error.response.data.message)
                } else {
                    switchErrorHandling(error.response.data)
                }
            }
        }
    )

    // タスクの削除処理
    const deleteTaskMutation = useMutation(
        {
            mutationFn: async (id: number) => {
                const res = axios.delete(`${process.env.REACT_APP_API_URL}/tasks/${id}`)
                return res
            },
            onSuccess: (_, variables) => {
                // 既存のタスクを取得
                const previousTasks = queryClient.getQueryData<Task[]>(['tasks'])
                if (previousTasks) {
                    // 削除したタスクを除外する
                    queryClient.setQueryData<Task[]>(
                        ['tasks'],
                        previousTasks.filter((task) => task.id !== variables)
                    )
                }
                resetEditedTask()
            },
            onError: (error: any) => {
                if (error.response.data.message) {
                    switchErrorHandling(error.response.data.message)
                } else {
                    switchErrorHandling(error.response.data)
                }
            }
        }
    )

    return { createTaskMutation, updateTaskMutation, deleteTaskMutation }
}