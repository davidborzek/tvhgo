import { LargeArrowLeftIcon, LargeArrowRightIcon } from '@/assets';
import { c } from '@/utils/classNames';
import styles from './GuideNavigation.module.scss';

type Props = {
  type?: 'left' | 'right';
  onClick: () => void;
};

function GuideNavigation({ type, onClick }: Props) {
  if (type === 'left') {
    return (
      <LargeArrowLeftIcon
        tabIndex={0}
        className={c(styles.left, styles.navigation)}
        onClick={onClick}
      />
    );
  }

  return (
    <LargeArrowRightIcon
      tabIndex={0}
      className={c(styles.right, styles.navigation)}
      onClick={onClick}
    />
  );
}

export default GuideNavigation;
