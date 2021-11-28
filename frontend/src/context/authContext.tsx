import { useState, ReactNode, Dispatch, SetStateAction } from 'react';
import { createCtx } from './utils';

const [useAuth, AuthContextProvider] = createCtx<{
  isAuthenticated: boolean;
  dispatch: Dispatch<SetStateAction<boolean>>;
}>();

export { useAuth };

export type AuthProviderProps = {
  children: ReactNode;
};

export const AuthProvider = ({ children }: AuthProviderProps) => {
  const [isAuthenticated, dispatch] = useState<boolean>(false);

  return (
    <AuthContextProvider value={{ isAuthenticated, dispatch }}>
      {children}
    </AuthContextProvider>
  );
};
