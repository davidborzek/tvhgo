import { useTranslation } from 'react-i18next';
import PaginationControlButton from './PaginationControlButton';
import styles from './PaginationControls.module.scss';

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
          props.onPageChange && props.onPageChange();
        }}
        label={'<<'}
        testID="first_page"
      />
      <PaginationControlButton
        disabled={props.limit > props.offset}
        onClick={() => {
          props.onPreviousPage();
          props.onPageChange && props.onPageChange();
        }}
        label={'<'}
        testID="previous_page"
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
          props.onPageChange && props.onPageChange();
        }}
        label={'>'}
        testID="next_page"
      />
      <PaginationControlButton
        disabled={props.offset + props.limit > props.total}
        onClick={() => {
          props.onLastPage();
          props.onPageChange && props.onPageChange();
        }}
        label={'>>'}
        testID="last_page"
      />
    </div>
  );
};

export default PaginationControls;
