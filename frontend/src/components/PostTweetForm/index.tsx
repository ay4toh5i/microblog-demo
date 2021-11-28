import { Form } from './Form';
import { usePostTweet } from './hooks';

export const PostTweetForm = () => {
  const { post } = usePostTweet();

  return (
    <Form onSubmit={post} />
  );
};
