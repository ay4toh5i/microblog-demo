import { useMutation } from 'react-query';
import type { AxiosResponse, AxiosError } from 'axios';
import { axiosApi } from 'http/axios';
import { useAuth } from 'context/authContext';

export type SignUpParams = {
  username: string;
  password: string;
  displayName: string;
  description: string;
};

export type Options = {
  onSuccess?: () => void;
};

export const useSignUp = (options?: Options) => {
  const { dispatch } = useAuth();

  const { mutateAsync, isLoading, error } = useMutation<
    AxiosResponse,
    AxiosError,
    SignUpParams
  >(
    params => {
      return axiosApi.post(
        '/users',
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
    signUp: mutateAsync,
    isLoading,
    error: error,
  };
};
