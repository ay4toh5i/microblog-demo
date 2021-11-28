import styled from 'styled-components';

export const Header = styled.header`
  display: flex;
  flex-direction: row;
  border-bottom: 1px solid var(--color-border);
  height: 3rem;
`;

export const HeaderItem = styled.button<{ active?: boolean }>`
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  font-weight: bold;
  background-color: ${props => props.active ? 'var(--color-background)' : '#EAEAEA'};
  border:none;

  :hover {
    cursor: pointer;
  }

  & + & {
    border-left: 1px solid var(--color-border);
  }
`;

