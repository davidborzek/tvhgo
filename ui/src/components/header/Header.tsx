import { BurgerMenuIcon, TvhgoHorizontalLogo } from '@/assets';

import styles from './Header.module.scss';

type Props = {
  onToggle: () => void;
};

function Header({ onToggle }: Props) {
  return (
    <div className={styles.root}>
      <div className={styles.bar}>
        <BurgerMenuIcon
          className={styles.menuIcon}
          onClick={() => onToggle()}
        />
        <TvhgoHorizontalLogo className={styles.logo} />
      </div>
    </div>
  );
}

export default Header;
