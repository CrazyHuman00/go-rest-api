import { FormEvent } from 'react'
import { useQueryClient } from '@tanstack/react-query'
import {
    ArrowRightOnRectangleIcon,
    ShieldCheckIcon,
} from '@heroicons/react/24/solid'
import useStore from '../store'
import { useQueryTasks } from '../hooks/useQueryTasks'
import { useMutateTask } from '../hooks/useMutateTask'
import { useMutateAuth } from '../hooks/useMutateAuth'
import { TaskItem } from './TaskItem'

export const Todo = () => {
    const queryClient = useQueryClient()
    const { editedTask } = useStore()
    const updateEditedTask = useStore((state) => state.updateEditedTask)
    const { data, isLoading } = useQueryTasks()
    const { createTaskMutation, updateTaskMutation} = useMutateTask()
    const { logoutMutation } = useMutateAuth()

    // サブミット処理
    const submitTaskHandler = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        if (!editedTask.title) return
        // 新規作成の場合（0の場合）はcreateTaskMutationを実行
        // 既存のタスクを更新する場合はupdateTaskMutationを実行
        if (editedTask.id === 0) {
            createTaskMutation.mutate({ title: editedTask.title })
        } else {
            updateTaskMutation.mutate(editedTask)
        }
    }

    // ログアウト処理
    const logout = async () => {
        await logoutMutation.mutateAsync()
        // キャッシュのクエリを削除
        queryClient.removeQueries({ queryKey: ['tasks'] })
    }

    return (
        <div className='flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono'>
            <div className='flex items-center my-3'>
                <ShieldCheckIcon className='h-8 w-8 mr-3 text-indigo-500 cursor-pointer' />
                <span className='text-center text-3xl font-extrabold'>
                    Task Manager
                </span>
            </div>
            <ArrowRightOnRectangleIcon
                onClick={logout}
                className="h-8 w-8 mr-6 text-blue-500"
            />
            <form onSubmit={submitTaskHandler}>
                <input
                    className='mb-3 mr-3 px-3 py-2 border border-gray-300'
                    placeholder='title ?'
                    type='text'
                    onChange={(e) => updateEditedTask({ ...editedTask, title: e.target.value })}
                    value={editedTask.title || ''}
                />
                <button
                    className='disabled:opacity-40 mx-3 px-3 py-3 text-white bg-indigo-600 rounded'
                    disabled={!editedTask.title}
                >
                    {editedTask.id === 0 ? 'Create' : 'Update'}
                </button>
            </form>

            {isLoading ? (
                <p>Loading...</p>
            ) : (
                <ul className='my-5'>
                    {data?.map((task) => (
                        <TaskItem key={task.id} id={task.id} title={task.title} />
                    ))}
                </ul>
            )}
        </div>
    )
}
