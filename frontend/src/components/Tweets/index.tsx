import { useTweets } from './hooks';
import { Tweet } from './Tweet';

export const Tweets = () => {
  const { tweets } = useTweets();

  return (
    <div>
      {tweets?.map(tweet => <Tweet key={tweet.id} tweet={tweet} />)}
    </div>
  );
};
