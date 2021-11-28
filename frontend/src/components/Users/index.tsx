import { User } from './User';
import { useUsers } from './hooks';

export const Users = () => {
  const { users } = useUsers();

  return (
    <div>{users?.map(user => <User key={user.id} user={user} />)}</div>
  );
};
