import styled from 'styled-components';
import { useForm } from 'react-hook-form';
import { Button } from 'components/Button';
import { FormItem } from 'components/FormItem';
import { Label } from 'components/Label';
import { Input } from 'components/Input';
import { useEffect } from 'react';

const StyledForm = styled.form`
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  width: 100%;
  padding: 2rem;
  margin-inline: auto;
  border-bottom: 1px solid var(--color-border);
`;

type FormData = {
  message: string,
};

export type Props = {
  onSubmit: (data: FormData) => Promise<any>,
};

export const Form = ({ onSubmit }: Props) => {
  const {
    register,
    handleSubmit,
    reset,
    formState: { isValid, isSubmitting, isSubmitSuccessful },
  } = useForm<FormData>({
    mode: 'onChange',
  });

  useEffect(() => {
    if (isSubmitSuccessful) {
      reset();
    }
  }, [isSubmitSuccessful]);

  return (
    <StyledForm onSubmit={handleSubmit(onSubmit)} noValidate>
      <FormItem>
        <Label htmlFor="message">post your message</Label>
        <Input 
          id="message" 
          type="text" 
          {...register('message', {
            required: true,
          })}
        />
      </FormItem>
      <Button type="submit" disabled={!isValid || isSubmitting}>post</Button>
    </StyledForm>
  );
};
