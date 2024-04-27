import { EpgEvent, ListResponse } from '@/clients/api/api.types';
import {
  LoaderFunctionArgs,
  Navigate,
  useLoaderData,
  useNavigate,
  useSearchParams,
} from 'react-router-dom';

import GuideEvent from '@/components/epg/guide/event/GuideEvent';
import PaginationControls from '@/components/common/paginationControls/PaginationControls';
import { getEpgEvents } from '@/clients/api/api';
import styles from './SearchView.module.scss';
import { usePagination } from '@/hooks/pagination';
import { useTranslation } from 'react-i18next';

const defaultLimit = 50;

export async function loader({ request }: LoaderFunctionArgs) {
  const urlParams = new URL(request.url).searchParams;
  const query = urlParams.get('q');

  if (!query) {
    return null;
  }

  return getEpgEvents({
    title: query,
    fullText: true,
    limit: defaultLimit,
    offset: parseInt(urlParams.get('offset') || '0') || 0,
  });
}

export const Component = () => {
  const data = useLoaderData() as ListResponse<EpgEvent> | null;
  const [searchParams, setSearchParams] = useSearchParams();
  const navigate = useNavigate();
  const { t } = useTranslation();

  const { limit, nextPage, previousPage, getOffset, firstPage, lastPage } =
    usePagination(defaultLimit, searchParams, setSearchParams);

  if (!data) {
    return <Navigate to={'/channels'} replace />;
  }

  return (
    <div className={styles.view}>
      <h1>
        {t('search_results_for', {
          query: searchParams.get('q'),
        })}
      </h1>

      {data.entries.length === 0 && <p>{t('no_results_found')}</p>}

      {data.entries.map((event) => (
        <GuideEvent
          key={event.id}
          eventId={event.id}
          title={event.title}
          subtitle={event.subtitle}
          description={event.description}
          endsAt={event.endsAt}
          startsAt={event.startsAt}
          onClick={(id) => {
            navigate(`/guide/events/${id}`);
          }}
          dvrState={event.dvrState}
          channel={event.channelName}
          showProgress
          showDate
        />
      ))}

      <PaginationControls
        onNextPage={nextPage}
        onPreviousPage={previousPage}
        onFirstPage={firstPage}
        onLastPage={() => lastPage(data.total)}
        limit={limit}
        offset={getOffset()}
        total={data.total}
      />
    </div>
  );
};

Component.displayName = 'SearchView';
