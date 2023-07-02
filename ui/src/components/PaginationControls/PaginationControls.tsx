import { useTranslation } from 'react-i18next';
import PaginationControlButton from './PaginationControlButton';
import styles from './PaginationControls.module.scss';

type Props = {
  onNextPage: () => void;
  onPreviousPage: () => void;
  onFirstPage: () => void;
  onLastPage: () => void;
  scrollTop?: () => void;
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

  return (
    <div className={styles.controls}>
      <PaginationControlButton
        disabled={props.limit > props.offset}
        onClick={() => {
          props.onFirstPage();
          props.scrollTop && props.scrollTop();
        }}
        label={'<<'}
      />
      <PaginationControlButton
        disabled={props.limit > props.offset}
        onClick={() => {
          props.onPreviousPage();
          props.scrollTop && props.scrollTop();
        }}
        label={'<'}
      />
      <span className={styles.page}>
        {t('pagination_info', {
          from: props.offset + 1,
          to: getMaxPageEntries(),
          total: props.total,
        })}
      </span>
      <PaginationControlButton
        disabled={props.offset + props.limit > props.total}
        onClick={() => {
          props.onNextPage();
          props.scrollTop && props.scrollTop();
        }}
        label={'>'}
      />
      <PaginationControlButton
        disabled={props.offset + props.limit > props.total}
        onClick={() => {
          props.onLastPage();
          props.scrollTop && props.scrollTop();
        }}
        label={'>>'}
      />
    </div>
  );
};

export default PaginationControls;
