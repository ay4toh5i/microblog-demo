import styled from 'styled-components';
import { useForm } from 'react-hook-form';
import { Button } from 'components/Button';
import { FormItem } from 'components/FormItem';
import { Label } from 'components/Label';
import { Input } from 'components/Input';

const StyledForm = styled.form`
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  width: 100%;
  max-width: 37rem;
  padding: 2rem;
  margin-inline: auto;
`;

type FormData = {
  username: string,
  password: string,
};

export type Props = {
  onSubmit: (data: FormData) => Promise<any>,
};

export const Form = ({ onSubmit }: Props) => {
  const {
    register,
    handleSubmit,
    formState: { isValid, isSubmitting },
  } = useForm<FormData>({
    mode: 'onChange',
  });

  return (
    <StyledForm onSubmit={handleSubmit(onSubmit)} noValidate>
      <h2>Login</h2>
      <FormItem>
        <Label htmlFor="username">username</Label>
        <Input 
          id="username" 
          type="text" 
          autoComplete="username" 
          {...register('username', {
            required: true,
          })}
        />
      </FormItem>
      <FormItem>
        <Label htmlFor="password">password</Label>
        <Input 
          id="password" 
          type="password" 
          autoComplete="current-password" 
          {...register('password', {
            required: true,
          })}
        />
      </FormItem>
      <Button type="submit" disabled={!isValid || isSubmitting}>submit</Button>
    </StyledForm>
  );
};
