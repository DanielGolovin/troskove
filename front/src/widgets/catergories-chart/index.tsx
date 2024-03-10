import { CategoriesChartItem } from "./ui/categories-chart-item";
import styles from "./index.module.css";
import { useState } from "react";
import { Category } from "src/entities/categories";
import { TransactionType } from "src/entities/transactions";
import { useGetCategoriesWithTotalQuery } from "./data-hooks/use-get-categories-with-total";

type CategoriesChartProps = {
  transactionType: TransactionType;
  onExludedCategoriesChange?: (id: Category["id"][]) => void;
};

export const CategoriesChart: React.FC<CategoriesChartProps> = ({
  onExludedCategoriesChange,
  transactionType,
}) => {
  const [excludedCategoryIds, setExcludedCategoryIds] = useState<
    Category["id"][]
  >([]);

  const { data = [] } = useGetCategoriesWithTotalQuery(transactionType);

  const toggleCategory = (id: Category["id"]) => {
    const newExcludedCategoryIds: Category["id"][] = [];

    if (excludedCategoryIds.includes(id)) {
      newExcludedCategoryIds.push(
        ...excludedCategoryIds.filter((i) => i !== id)
      );
    } else {
      newExcludedCategoryIds.push(id, ...excludedCategoryIds);
    }

    setExcludedCategoryIds(newExcludedCategoryIds);
    onExludedCategoriesChange?.(newExcludedCategoryIds);
  };

  const total = data.reduce((acc, item) => acc + item.total, 0) || 0;

  return (
    <div className={styles.container}>
      {data?.map((item) => (
        <CategoriesChartItem
          name={item.name}
          total={item.total}
          onClick={() => toggleCategory(item.id)}
          relativeValue={(item.total / total) * 100}
          disabled={excludedCategoryIds.includes(item.id)}
          key={item.id}
        />
      ))}
    </div>
  );
};
