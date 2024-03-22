import { EpgEvent } from '@/clients/api/api.types';
import { cleanup, render } from '@testing-library/react';
import { useLoaderData, useNavigate, useSearchParams } from 'react-router-dom';
import { afterEach, beforeEach, expect, test, vi, describe } from 'vitest';
import { Component as ChannelListView } from './ChannelListView';
import userEvent from '@testing-library/user-event';
import { TestIds } from '@/__test__/ids';
import { usePagination } from '@/hooks/pagination';

vi.mock('react-router-dom');
vi.mock('@/hooks/pagination');

const buildChannel = (id: string, name: string): EpgEvent => {
  return {
    audioDesc: true,
    channelId: id,
    channelName: name,
    channelNumber: 1,
    description: `${name} - Description`,
    endsAt: 0,
    hd: true,
    id: 1,
    nextEventId: 2,
    piconId: 1,
    startsAt: 0,
    subtitle: `${name} - Subtitle`,
    subtitled: true,
    title: `${name} - Title`,
    widescreen: true,
  };
};

const buildChannels = (count: number): EpgEvent[] => {
  return [...new Array(count)].map((_, i) =>
    buildChannel(`id_${i}`, `Channel_${i}`)
  );
};

const navigateMock = vi.fn();
const setQueryParamsMock = vi.fn();
const nextPageMock = vi.fn();
const previousPageMock = vi.fn();
const getOffsetMock = vi.fn();
const firstPageMock = vi.fn();
const lastPageMock = vi.fn();

beforeEach(() => {
  vi.mocked(useSearchParams).mockReturnValue([
    new URLSearchParams(),
    setQueryParamsMock,
  ]);
  vi.mocked(useNavigate).mockReturnValue(navigateMock);
  vi.mocked(usePagination).mockReturnValue({
    firstPage: firstPageMock,
    getOffset: getOffsetMock,
    lastPage: lastPageMock,
    limit: 50,
    nextPage: nextPageMock,
    previousPage: previousPageMock,
    setLimit: vi.fn(),
  });
});

afterEach(() => {
  vi.resetAllMocks();
  cleanup();
});

test('render empty state', () => {
  vi.mocked(useLoaderData).mockReturnValue({
    entries: [],
    offset: 0,
    total: 0,
  });
  getOffsetMock.mockReturnValue(0);

  const document = render(<ChannelListView />);

  expect(document.asFragment()).toMatchSnapshot();
});

test('render with one page', () => {
  const channels = buildChannels(5);
  vi.mocked(useLoaderData).mockReturnValue({
    entries: channels,
    offset: 0,
    total: channels.length,
  });
  getOffsetMock.mockReturnValue(0);

  const document = render(<ChannelListView />);

  expect(document.asFragment()).toMatchSnapshot();
});

test('render with multiple pages', () => {
  const channels = buildChannels(5);
  vi.mocked(useLoaderData).mockReturnValue({
    entries: channels,
    offset: 0,
    total: 100,
  });
  getOffsetMock.mockReturnValue(0);

  const document = render(<ChannelListView />);

  expect(document.asFragment()).toMatchSnapshot();
});

describe('pagination', () => {
  test('go next page', async () => {
    const channels = buildChannels(5);
    vi.mocked(useLoaderData).mockReturnValue({
      entries: channels,
      offset: 0,
      total: 200,
    });
    getOffsetMock.mockReturnValue(0);

    const document = render(<ChannelListView />);

    await userEvent.click(document.getByTestId(TestIds.PAGINATION_NEXT_PAGE));
    expect(nextPageMock).toHaveBeenCalled();
  });

  test('go previous', async () => {
    const channels = buildChannels(5);
    vi.mocked(useLoaderData).mockReturnValue({
      entries: channels,
      offset: 50,
      total: 200,
    });
    getOffsetMock.mockReturnValue(50);

    const document = render(<ChannelListView />);

    await userEvent.click(
      document.getByTestId(TestIds.PAGINATION_PREVIOUS_PAGE)
    );
    expect(previousPageMock).toHaveBeenCalled();
  });

  test('go last', async () => {
    const channels = buildChannels(5);
    vi.mocked(useLoaderData).mockReturnValue({
      entries: channels,
      offset: 0,
      total: 200,
    });
    getOffsetMock.mockReturnValue(0);

    const document = render(<ChannelListView />);

    await userEvent.click(document.getByTestId(TestIds.PAGINATION_LAST_PAGE));
    expect(lastPageMock).toHaveBeenCalled();
  });

  test('go first', async () => {
    const channels = buildChannels(5);
    vi.mocked(useLoaderData).mockReturnValue({
      entries: channels,
      offset: 200,
      total: 200,
    });
    getOffsetMock.mockReturnValue(200);

    const document = render(<ChannelListView />);

    await userEvent.click(document.getByTestId(TestIds.PAGINATION_FIRST_PAGE));
    expect(firstPageMock).toHaveBeenCalled();
  });
});
