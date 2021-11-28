import styled from 'styled-components';
import { FollowButton } from './FollowButton';
import type { User as UserData } from './types';
import { useFollow } from './hooks';

const Container = styled.div`
  display: flex;
  flex-direction: column;
  padding: 1rem;

  & + & {
    border-top: 1px solid var(--color-border)
  }
`;

const Header = styled.header`
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
`;

const Identity = styled.div`
  display: flex;
  flex-direction: column;

  .displayName {
    font-weight: bold;
  }
`;

type Props = {
  user: UserData;
};

export const User = ({ user }: Props) => {
  const { follow } = useFollow();

  const followHandler = () => {
    follow({ userId: user.id });
  };

  return (
    <Container>
      <Header>
        <Identity>
          <div className="displayName">{user.displayName}</div>
          <div>@{user.username} {user.followedBy && 'followed you'}</div>
        </Identity>
        <FollowButton following={user.following} onClick={followHandler}/>
      </Header>
      <p>{user.description}</p>
    </Container>
  );
};
