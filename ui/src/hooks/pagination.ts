import { useState } from 'react';
import { SetURLSearchParams } from 'react-router-dom';
import { URLSearchParams } from 'url';

export const usePagination = (
  initialLimit: number = 50,
  searchParams: URLSearchParams,
  setSearchParams: SetURLSearchParams
) => {
  const [limit, setLimit] = useState(initialLimit);

  const getOffset = () => {
    const offset = searchParams.get('offset');
    return offset ? parseInt(offset) : 0;
  };

  const setOffset = (value: number) => {
    setSearchParams((prev) => {
      prev.set('offset', `${value}`);
      return prev;
    });
  };

  const nextPage = () => {
    setOffset(getOffset() + limit);
  };

  const previousPage = () => {
    setOffset(getOffset() - limit);
  };

  const firstPage = () => {
    setOffset(0);
  };

  const lastPage = (total: number) => {
    setOffset(Math.floor(total / limit) * limit);
  };

  return {
    limit,
    nextPage,
    previousPage,
    getOffset,
    lastPage,
    firstPage,
    setLimit,
  };
};
