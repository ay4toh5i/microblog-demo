import { createGlobalStyle } from 'styled-components';

export const GlobalStyles = createGlobalStyle`
  :root {
    --color-main: #094067;
    --color-secondary: #90b4ce;
    --color-text: #5f6c7b;
    --color-highlight: #3da9fc;
    --color-outline-text: #fffffe;
    --color-background: #fffffe;
    --color-border: lightgray;
  }

  body {
    font-family: "ヒラギノ角ゴ Pro W3", "Hiragino Kaku Gothic Pro", "メイリオ", "Meiryo", sans-serif;
    background-color: var(--color-background);
    color: var(--color-text);
  }
`;
