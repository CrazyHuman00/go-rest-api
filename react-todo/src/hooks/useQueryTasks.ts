import axios from "axios"
import { useQuery, UseQueryOptions } from "@tanstack/react-query"
import { Task } from "../types"
import { useError } from "../hooks/useError"

export const useQueryTasks = () => {
    const { switchErrorHandling } = useError()
    const getTasks = async () => {
        const { data } = await axios.get<Task[]>(
            `${process.env.REACT_APP_API_URL}/tasks`,
            { withCredentials: true }
        )
        return data
    }

    // fetchで取得したデータをクライアントに格納
    return useQuery<Task[], Error>({
        queryKey: ['tasks'],
        queryFn: getTasks,
        staleTime: Infinity, // キャッシュの有効期限を無期限に設定
        onError: (error: any) => {
            if (error.response.data.message) {
                switchErrorHandling(error.response.data.message)
            } else {
                switchErrorHandling(error.response.data)
            }
        },
    } as UseQueryOptions<Task[], Error>)
}