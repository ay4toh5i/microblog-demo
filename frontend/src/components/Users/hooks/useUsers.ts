import { useQuery } from 'react-query';
import { axiosApi } from 'http/axios';
import type { User } from '../types';

type Response = {
  users: User[];
};

export const useUsers = () => {
  const { data, isLoading } = useQuery<User[]>(
    ['users'],
    async () => {
      const { data } = await axiosApi.get<Response>('/users' , { withCredentials: true });

      return data.users;
    },
  );

  return { users: data, isLoading };
};
