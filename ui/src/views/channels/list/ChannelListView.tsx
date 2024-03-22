import {
  LoaderFunctionArgs,
  useLoaderData,
  useNavigate,
  useSearchParams,
} from 'react-router-dom';
import { useTranslation } from 'react-i18next';

import ChannelListItem from '@/components/channels/listItem/ChannelListItem';
import { usePagination } from '@/hooks/pagination';
import PaginationControls from '@/components/common/paginationControls/PaginationControls';
import EmptyState from '@/components/common/emptyState/EmptyState';

import styles from './ChannelListView.module.scss';
import { getEpgEvents } from '@/clients/api/api';
import { EpgEvent, ListResponse } from '@/clients/api/api.types';

const defaultLimit = 50;

export async function loader({ request }: LoaderFunctionArgs) {
  const query = new URL(request.url).searchParams;

  return getEpgEvents({
    nowPlaying: true,
    limit: defaultLimit,
    offset: parseInt(query.get('offset') || '0') || 0,
    sort_key: 'channelNumber',
    sort_dir: 'asc',
  });
}

export function Component() {
  const { t } = useTranslation();
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();

  const { limit, nextPage, previousPage, getOffset, firstPage, lastPage } =
    usePagination(defaultLimit, searchParams, setSearchParams);

  const { entries, total } = useLoaderData() as ListResponse<EpgEvent>;

  if (entries.length === 0) {
    return (
      <EmptyState title={t('no_channels')} subtitle={t('no_channels_info')} />
    );
  }

  const renderChannels = () => {
    return entries.map((event) => {
      return (
        <ChannelListItem
          key={event.channelId}
          event={event}
          onClick={(id) => {
            navigate(`/channels/${id}`);
          }}
        />
      );
    });
  };

  return (
    <div className={styles.container}>
      <div className={styles.channelList}>{renderChannels()}</div>
      <PaginationControls
        onNextPage={nextPage}
        onPreviousPage={previousPage}
        onFirstPage={firstPage}
        onLastPage={() => lastPage(total)}
        limit={limit}
        offset={getOffset()}
        total={total}
      />
    </div>
  );
}

Component.displayName = 'ChannelListView';
