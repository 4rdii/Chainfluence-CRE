import { createBrowserRouter } from 'react-router';
import { AppLayout } from './components/layout/app-layout';
import { ConnectWallet } from './pages/connect-wallet';
import { Landing } from './pages/landing';
import { Dashboard } from './pages/dashboard';
import { Marketplace } from './pages/marketplace';
import { CreateCampaign } from './pages/create-campaign';
import { CampaignView } from './pages/campaign-view';
import { ApplyCampaign } from './pages/apply-campaign';
import { ChannelView } from './pages/channel-view';
import { ProposeDeal } from './pages/propose-deal';
import { DealView } from './pages/deal-view';
import { MyCampaigns } from './pages/my-campaigns';
import { MyChannels } from './pages/my-channels';
import { MyDeals } from './pages/my-deals';
import { Settings } from './pages/settings';
import { Docs } from './pages/docs';

export const router = createBrowserRouter([
  {
    path: '/',
    element: <Landing />,
  },
  {
    path: '/connect',
    element: <ConnectWallet />,
  },
  {
    path: '/',
    element: <AppLayout />,
    children: [
      {
        path: '/dashboard',
        element: <Dashboard />,
      },
      {
        path: '/marketplace',
        element: <Marketplace />,
      },
      {
        path: '/create-campaign',
        element: <CreateCampaign />,
      },
      {
        path: '/campaign/:id',
        element: <CampaignView />,
      },
      {
        path: '/apply/:id',
        element: <ApplyCampaign />,
      },
      {
        path: '/channel/:id',
        element: <ChannelView />,
      },
      {
        path: '/propose-deal/:channelId',
        element: <ProposeDeal />,
      },
      {
        path: '/deal/:id',
        element: <DealView />,
      },
      {
        path: '/my-campaigns',
        element: <MyCampaigns />,
      },
      {
        path: '/my-channels',
        element: <MyChannels />,
      },
      {
        path: '/my-deals',
        element: <MyDeals />,
      },
      {
        path: '/settings',
        element: <Settings />,
      },
      {
        path: '/docs',
        element: <Docs />,
      },
    ],
  },
]);
