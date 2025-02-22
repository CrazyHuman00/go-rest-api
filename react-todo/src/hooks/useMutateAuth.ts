import axios, { AxiosResponse } from 'axios'
import { useNavigate } from 'react-router-dom'
import { useMutation } from '@tanstack/react-query'
import useStore from '../store'
import { Credential } from '../types'
import { useError } from '../hooks/useError'


export const useMutateAuth = () => {
    const navigate = useNavigate()
    const resetEditedTask = useStore((state) => state.resetEditedTask)
    const { switchErrorHandling } = useError()

    // ログイン処理
    const loginMutation = useMutation(
        {
            // emailとpasswordを引数に取る
            mutationFn: async (user: Credential) => {
                const res = await axios.post(`${process.env.REACT_APP_API_URL}/login`, user)
                return res
            },
            onSuccess: (res: AxiosResponse) => {
                navigate('/todo')
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

    // サインアップ処理
    const registerMutation = useMutation(
        {
            // emailとpasswordを引数に取る
            mutationFn: async (user: Credential) => {
                const res = await axios.post(`${process.env.REACT_APP_API_URL}/signup`, user)
                return res
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

    // ログアウト処理
    const logoutMutation = useMutation(
        {
            // ログアウト処理を行う
            mutationFn: async () => {
                await axios.post(`${process.env.REACT_APP_API_URL}/logout`)
            },
            onSuccess: () => {
                resetEditedTask()
                navigate('/')
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

    return { loginMutation, registerMutation, logoutMutation }
}