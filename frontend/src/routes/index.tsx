import { useRoutes } from 'react-router-dom';

import { Landing } from '@/features/misc';

export const AppRoutes = () => {

  const commonRoutes = [{ path: '/', element: <Landing /> }];

  // const routes = auth.user ? protectedRoutes : publicRoutes;

  // const element = useRoutes([...routes, ...commonRoutes]);
  const element = useRoutes(commonRoutes);

  return <>{element}</>;
};
