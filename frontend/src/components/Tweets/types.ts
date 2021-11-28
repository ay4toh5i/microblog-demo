export type Tweet = {
  id: string;
  user: {
    id: number;
    username: string;
    displayName: string;
    description: string;
  };
  message: string;
};

