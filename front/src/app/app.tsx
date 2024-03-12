import { useState } from "react";
import { CategoriesChart } from "../widgets/catergories-chart";
import { TransactionTypeSelector } from "../widgets/transaction-type-selector";
import { Category } from "../entities/categories";

import { TransactionType } from "../entities/transactions";
import { TransactionsList } from "src/widgets/transactions-list";

import "./app.css";

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
      <TransactionsList />
    </div>
  );
}

export default App;
