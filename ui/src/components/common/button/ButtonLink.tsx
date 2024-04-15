import { ButtonStyle, getStyleClass } from './Button';

import { c } from '@/utils/classNames';
import styles from './Button.module.scss';

type Props = {
  label: string;
  download?: string | boolean;
  size?: 'small' | 'medium' | 'large';
  href?: string;
  quiet?: boolean;
  className?: string;
  style?: ButtonStyle;
};

const ButtonLink = (props: Props) => {
  return (
    <a
      className={c(
        styles.button,
        props.size ? styles[props.size] : styles.medium,
        props.quiet ? styles.quiet : '',
        getStyleClass(props.style),
        props.className ? props.className : ''
      )}
      href={props.href}
      download={props.download}
    >
      {props.label}
    </a>
  );
};

export default ButtonLink;
