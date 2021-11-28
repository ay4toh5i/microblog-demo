import { useState } from 'react';
import { PostTweetForm } from 'components/PostTweetForm';
import { Tweets } from 'components/Tweets';
import { Users } from 'components/Users';
import { Header, HeaderItem } from 'components/Header';

export const Timeline = () => {
  const [showTweets, setShowTweets] = useState(true);

  return (
    <main>
      <Header>
        <HeaderItem active={showTweets} onClick={() => setShowTweets(true)}>tweets</HeaderItem>
        <HeaderItem active={!showTweets} onClick={() => setShowTweets(false)}>users</HeaderItem>
      </Header>
      {showTweets ? (
        <div>
          <PostTweetForm />
          <Tweets />
        </div>
      ) : <Users />}
    </main>
  );
};
