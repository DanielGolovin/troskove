import { useQuery } from '@tanstack/react-query'
import { CategoryWithTotal } from '../../../entities/categories';
import { TransactionType } from 'src/entities/transactions';

export const getCategoriesQueryKeyBase = 'categories'

export const useGetCategoriesWithTotalQuery = (transactionType: TransactionType) => {
  const { data } = useQuery<CategoryWithTotal[]>({
    queryKey: [getCategoriesQueryKeyBase, transactionType],
    queryFn: () => Promise.resolve(mockData[transactionType]),
    // queryFn: () => fetch('/api/categories').then((res) => res.json()),
  });

  return { data };
}


const mockDataExpenses: CategoryWithTotal[] = [
    {
      id: "1",
      name: "Transport",
      total: -100,
    },
    {
      id: "2",
      name: "Rent",
      total: -900,
    },
    {
      id: "3",
      name: "Food",
      total: -400,
    },
  ].sort((a, b) => Math.abs(b.total) - Math.abs(a.total));

const mockDataIncome: CategoryWithTotal[] = [
    {
      id: "1",
      name: "Salary",
      total: 1000,
    },
    {
      id: "2",
      name: "Freelance",
      total: 200,
    },
  ].sort((a, b) => Math.abs(b.total) - Math.abs(a.total));

const mockDataMixed: CategoryWithTotal[] = [
    {
      id: "1",
      name: "Transport",
      total: 100,
    },
    {
      id: "2",
      name: "Rent",
      total: 900,
    },
    {
      id: "3",
      name: "Food",
      total: 400,
    },
    {
      id: "4",
      name: "Salary",
      total: 1000,
    },
    {
      id: "5",
      name: "Freelance",
      total: 200,
    },
  ].sort((a, b) => Math.abs(b.total) - Math.abs(a.total));

  const mockData = {
    expenses: mockDataExpenses,
    incomes: mockDataIncome,
    mixed: mockDataMixed,
  };
  