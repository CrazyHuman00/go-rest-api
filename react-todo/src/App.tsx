// Appコンポーネントは、アプリケーションの起動時に実行されるコンポーネント。

import { useEffect } from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { Auth } from './components/Auth';
import { Todo } from './components/Todo';
import axios from 'axios';
import { CsrfToken } from './types';

function App() {
  useEffect(() => {
    axios.defaults.withCredentials = true;
    // 関数を定義して即時実行
    const getCsrfToken = async () => {
      // CSRFトークンを取得
      const { data } = await axios.get<CsrfToken>(
        `${process.env.REACT_APP_API_URL}/csrf`
      )
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    getCsrfToken()
  }, [])

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Auth />} />
        <Route path="/todo" element={<Todo />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
