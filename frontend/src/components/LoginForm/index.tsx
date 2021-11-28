import { Form } from './Form';
import { useLogin } from './hooks';

export const LoginForm = () => {
  const { login } = useLogin();

  return (
    <Form onSubmit={data => login(data)}/>
  );
};
