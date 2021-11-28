import styled from 'styled-components';

export const Button = styled.button`
  height: 2.5rem;
  background-color: var(--color-highlight);
  color: var(--color-outline-text);
  font-weight: bold;
  border: none;
  border-radius: 0.5rem;

  :hover {
    cursor: pointer
  }

  :disabled {
    background-color: var(--color-secondary);
  }
`;
