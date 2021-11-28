import React from 'react'
import ReactDOM from 'react-dom'
import { QueryClient, QueryClientProvider } from 'react-query';
import { GlobalStyles } from 'styles/GlobalStyles';
import 'modern-css-reset/dist/reset.min.css';
import { AuthProvider } from 'context/authContext';
import { App } from './App' 

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      notifyOnChangeProps: 'tracked',
    },
  },
});

ReactDOM.render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <App />
      </AuthProvider>
    </QueryClientProvider>
    <GlobalStyles />
  </React.StrictMode>,
  document.getElementById('root')
)
