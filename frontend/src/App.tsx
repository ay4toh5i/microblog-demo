import { Layout } from 'components/Layout';
import { useAuth } from 'context/authContext';
import { Top } from 'pages/Top';
import { Timeline } from 'pages/Timeline';

export const  App = () => {
  const { isAuthenticated } = useAuth();

  if (isAuthenticated) {
    return (
      <Layout>
        <Timeline />
      </Layout>
    );
  }

  return (
    <Layout>
      <Top />
    </Layout>
  )
};
