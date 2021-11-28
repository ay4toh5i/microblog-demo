import { useQuery } from 'react-query';
import { axiosApi } from 'http/axios';
import type { Tweet } from './types';

type Response = {
  tweets: Tweet[];
};

export const useTweets = () => {
  const { data, isLoading } = useQuery<Tweet[]>(
    ['tweets'],
    async () => {
      const { data } = await axiosApi.get<Response>('/tweets' , { withCredentials: true });

      return data.tweets;
    },
  );

  return { tweets: data, isLoading };
};
