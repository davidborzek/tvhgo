import PaginationControlButton from './PaginationControlButton';
import { TestIds } from '@/__test__/ids';
import styles from './PaginationControls.module.scss';
import { useTranslation } from 'react-i18next';

type Props = {
  onNextPage: () => void;
  onPreviousPage: () => void;
  onFirstPage: () => void;
  onLastPage: () => void;
  onPageChange?: () => void;
  limit: number;
  offset: number;
  total: number;
};

const PaginationControls = (props: Props) => {
  const { t } = useTranslation();
  const getMaxPageEntries = () => {
    const next = props.offset + props.limit;
    return next > props.total ? props.total : next;
  };

  if (props.total === 0) {
    return <></>;
  }

  return (
    <div className={styles.controls}>
      <PaginationControlButton
        disabled={props.limit > props.offset}
        onClick={() => {
          props.onFirstPage();
          if (props.onPageChange) props.onPageChange();
        }}
        label={'<<'}
        testID={TestIds.PAGINATION_FIRST_PAGE}
      />
      <PaginationControlButton
        disabled={props.limit > props.offset}
        onClick={() => {
          props.onPreviousPage();
          if (props.onPageChange) props.onPageChange();
        }}
        label={'<'}
        testID={TestIds.PAGINATION_PREVIOUS_PAGE}
      />
      <span className={styles.page}>
        {t('pagination_info', {
          from: props.total > 0 ? props.offset + 1 : 0,
          to: getMaxPageEntries(),
          total: props.total,
        })}
      </span>
      <PaginationControlButton
        disabled={props.offset + props.limit > props.total}
        onClick={() => {
          props.onNextPage();
          if (props.onPageChange) props.onPageChange();
        }}
        label={'>'}
        testID={TestIds.PAGINATION_NEXT_PAGE}
      />
      <PaginationControlButton
        disabled={props.offset + props.limit > props.total}
        onClick={() => {
          props.onLastPage();
          if (props.onPageChange) props.onPageChange();
        }}
        label={'>>'}
        testID={TestIds.PAGINATION_LAST_PAGE}
      />
    </div>
  );
};

export default PaginationControls;
