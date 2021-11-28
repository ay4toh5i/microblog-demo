import { useMutation, useQueryClient } from 'react-query';
import type { AxiosResponse, AxiosError } from 'axios';
import { axiosApi } from 'http/axios';

export type LoginParams = {
  message: string;
};

export type Options = {
  onSuccess?: () => void;
};

export const usePostTweet = (options?: Options) => {
  const queryClient = useQueryClient();

  const { mutateAsync, isLoading, error } = useMutation<
    AxiosResponse,
    AxiosError,
    LoginParams
  >(
    params => {
      return axiosApi.post(
        '/tweets',
        params,
        {
          withCredentials: true,
        },
      );
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('tweets');
        options?.onSuccess?.();
      },
    },
  );

  return {
    post: mutateAsync,
    isLoading,
    error: error,
  };
};
