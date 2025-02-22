import { create } from 'zustand';

// 状態管理したいstateの型を定義
type EditedTask = {
    id: number
    title: string
}

type State = {
    editedTask: EditedTask
    updateEditedTask: (payload: EditedTask) => void
    resetEditedTask: () => void
}

// 初期状態を定義
const useStore = create<State>((set) => ({
    editedTask: { id: 0, title: '' },
    updateEditedTask: (payload) => set({ editedTask: payload, }),
    resetEditedTask: () => set({ editedTask: { id: 0, title: '' } }),
}))

export default useStore