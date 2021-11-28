import { useState } from 'react';
import { SignUpForm } from 'components/SignUpForm';
import { LoginForm } from 'components/LoginForm';
import { Header, HeaderItem } from 'components/Header';

export const Top = () => {
  const [showSignUp, setShowSignUp] = useState(true);

  return (
    <main>
      <Header>
        <HeaderItem active={showSignUp} onClick={() => setShowSignUp(true)}>SignUp</HeaderItem>
        <HeaderItem active={!showSignUp} onClick={() => setShowSignUp(false)}>Login</HeaderItem>
      </Header>
      {showSignUp ? <SignUpForm /> : <LoginForm />}
    </main>
  );
};
