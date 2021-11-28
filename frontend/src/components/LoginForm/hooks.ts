import { useMutation } from 'react-query';
import type { AxiosResponse, AxiosError } from 'axios';
import { axiosApi } from 'http/axios';
import { useAuth } from 'context/authContext';

export type LoginParams = {
  username: string;
  password: string;
};

export type Options = {
  onSuccess?: () => void;
};

export const useLogin = (options?: Options) => {
  const { dispatch } = useAuth();

  const { mutateAsync, isLoading, error } = useMutation<
    AxiosResponse,
    AxiosError,
    LoginParams
  >(
    params => {
      return axiosApi.post(
        '/login',
        params,
        {
          withCredentials: true,
        },
      );
    },
    {
      onSuccess: () => {
        dispatch(true);
        options?.onSuccess?.();
      },
    },
  );

  return {
    login: mutateAsync,
    isLoading,
    error: error,
  };
};
