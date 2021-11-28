import styled from 'styled-components';

const StyledButton = styled.button<{ following: boolean }>`
  height: 2rem;
  width: 10rem;
  border-radius: 0.5rem;;
  background-color: ${props => props.following ? 'var(--color-main)' : 'var(--color-secondary)'};
  color: var(--color-outline-text);
  font-weight: bold;
  border: none;

  :hover {
    cursor: pointer;
  }
`;

type Props = {
  following: boolean;
  onClick?: () => void;
};

export const FollowButton = ({ following, onClick }: Props) => {
  return (
    <StyledButton following={following} onClick={onClick}>
      {following ? 'unfollow' : 'follow'}
    </StyledButton>
  );
};
