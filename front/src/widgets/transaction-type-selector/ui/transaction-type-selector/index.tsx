import { TransactionType } from "src/entities/transactions";
import styles from "./index.module.css";

type TransactionsTypeSelectorProps = {
  total: string;
  selected: TransactionType;
  types: TransactionType[];
  onSelect: (type: TransactionType) => void;
};

export const TransactionsTypeSelector: React.FC<
  TransactionsTypeSelectorProps
> = ({ total, selected, types, onSelect }) => {
  const items = types.map((type) => {
    const classList = [styles["item"]];

    if (type === selected) {
      classList.push(styles["item-selected"]);
    }

    return (
      <div
        key={type}
        className={classList.join(" ")}
        onClick={() => onSelect(type)}
      >
        {type}
      </div>
    );
  });

  return (
    <div className={styles["container"]}>
      <nav className={styles["items-container"]}>{items}</nav>
      <div className={styles["total"]}>{total}</div>
    </div>
  );
};
