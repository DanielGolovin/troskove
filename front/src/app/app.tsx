import { useState } from "react";
import { CategoriesChart } from "../widgets/catergories-chart";
import { TransactionTypeSelector } from "../widgets/transaction-type-selector";
import { Category } from "../entities/categories";

import "./app.css";
import { TransactionType } from "../entities/transactions";

type AppModel = {
  selectedTransactionType: TransactionType;
  setSelectedTransactionType: (type: TransactionType) => void;
  excludedCategoryIds: Category["id"][];
  setExcludedCategoryIds: (categories: Category["id"][]) => void;
};

const useAppModel = (): AppModel => {
  const [selectedTransactionType, setSelectedTransactionType] =
    useState<TransactionType>("mixed");
  const [excludedCategoryIds, setExcludedCategoryIds] = useState<
    Category["id"][]
  >([]);

  return {
    selectedTransactionType,
    setSelectedTransactionType,
    excludedCategoryIds,
    setExcludedCategoryIds,
  };
};

function App() {
  const {
    selectedTransactionType,
    setSelectedTransactionType,
    excludedCategoryIds,
    setExcludedCategoryIds,
  } = useAppModel();

  return (
    <div>
      <div className={"container"}>
        <TransactionTypeSelector
          excludedCategoryIds={excludedCategoryIds}
          onTransactionTypeChange={setSelectedTransactionType}
        />
      </div>
      <div className={"container"}>
        <CategoriesChart
          transactionType={selectedTransactionType}
          onExludedCategoriesChange={setExcludedCategoryIds}
        />
      </div>
    </div>
  );
}

export default App;
