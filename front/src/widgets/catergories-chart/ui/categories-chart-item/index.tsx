import clsx from "clsx";
import styles from "./index.module.css";

type CategoriesChartItemProps = {
  name: string;
  total: number;
  relativeValue: number; // percentage
  onClick: () => void;
  disabled: boolean;
};

export const CategoriesChartItem: React.FC<CategoriesChartItemProps> = ({
  name,
  total,
  onClick,
  disabled,
  relativeValue,
}) => {
  return (
    <div
      className={clsx(styles.item, {
        [styles["item-disabled"]]: disabled,
      })}
      onClick={onClick}
    >
      <div
        style={{
          width: `${relativeValue}%`,
        }}
        className={styles.indicator}
      ></div>
      <div className={styles["item-data"]}>
        <div>{name}</div>
        <div>{total}$</div>
      </div>
    </div>
  );
};
