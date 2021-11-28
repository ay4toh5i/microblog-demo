import { Form } from './Form';
import { useSignUp } from './hooks';

export const SignUpForm = () => {
  const { signUp } = useSignUp();

  return (
    <Form onSubmit={data => signUp(data)}/>
  );
};
