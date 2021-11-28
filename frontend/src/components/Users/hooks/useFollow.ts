import { useMutation, useQueryClient } from 'react-query';
import type { AxiosResponse, AxiosError } from 'axios';
import { axiosApi } from 'http/axios';

export type FollowParams = {
  userId: number;
};

export type Options = {
  onSuccess?: () => void;
};

export const useFollow = (options?: Options) => {
  const queryClient = useQueryClient();

  const { mutateAsync, isLoading, error } = useMutation<
    AxiosResponse,
    AxiosError,
    FollowParams
  >(
    params => {
      return axiosApi.post(
        '/followings',
        params,
        {
          withCredentials: true,
        },
      );
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries('users');
        options?.onSuccess?.();
      },
    },
  );

  return {
    follow: mutateAsync,
    isLoading,
    error: error,
  };
};
