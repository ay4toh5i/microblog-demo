import styled from 'styled-components';
import type { Tweet as TweetData } from './types';

const Container = styled.article`
  display: flex;
  flex-direction: column;
  padding: 1rem;

  & + & {
    border-top: 1px solid var(--color-border);
  }
`;

const Header = styled.header`
  display: flex;
  flex-direction: row;
  gap: 0.5rem;

  .displayName {
    font-weight: bold;
  }
`;

type Props = {
  tweet: TweetData;
};

export const Tweet = ({ tweet: { message, user } }: Props) => (
  <Container>
    <Header>
      <div className="displayName">{user.displayName}</div>
      <div>@{user.username}</div>
    </Header>
    <p>{message}</p>
  </Container>
);
