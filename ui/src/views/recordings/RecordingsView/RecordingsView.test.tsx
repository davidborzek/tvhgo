import { useFetchRecordings, useManageRecordings } from '@/hooks/recording';
import { render, cleanup } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { afterEach, beforeEach, describe, expect, test, vi } from 'vitest';
import RecordingsView from './RecordingsView';
import { useLoading } from '@/contexts/LoadingContext';
import { Recording } from '@/clients/api/api.types';
import { TestIds } from '@/__test__/ids';

vi.mock('@/contexts/LoadingContext');
vi.mock('@/hooks/recording');
vi.mock('react-router-dom');

const navigateMock = vi.fn();
const stopAndCancelRecordingsMock = vi.fn();
const fetchMock = vi.fn();
const removeRecordingsMock = vi.fn();
const setQueryParamsMock = vi.fn();

beforeEach(() => {
  vi.mocked(useLoading).mockReturnValue({
    isLoading: false,
    setIsLoading: vi.fn(),
  });

  vi.mocked(useManageRecordings).mockReturnValue({
    pending: false,
    removeRecordings: removeRecordingsMock,
    stopAndCancelRecordings: stopAndCancelRecordingsMock,
  });

  vi.mocked(useFetchRecordings).mockReturnValue({
    recordings: [],
    error: null,
    fetch: fetchMock,
    total: 0,
  });

  vi.mocked(useSearchParams).mockReturnValue([
    new URLSearchParams(),
    setQueryParamsMock,
  ]);
  vi.mocked(useNavigate).mockReturnValue(navigateMock);

  vi.mocked(stopAndCancelRecordingsMock).mockResolvedValue(null);
  vi.mocked(removeRecordingsMock).mockResolvedValue(null);
});

afterEach(() => {
  vi.restoreAllMocks();
  cleanup();
});

test.each([
  [null, 'upcoming'],
  ['upcoming', 'upcoming'],
  ['finished', 'finished'],
  ['failed', 'failed'],
  ['removed', 'removed'],
])(
  'should render error: status=%s',
  (status: string | null, expectedStatus: string) => {
    mockStatus(status);
    vi.mocked(useFetchRecordings).mockReturnValue({
      recordings: [],
      error: 'some error occurred',
      fetch: vi.fn(),
      total: 0,
    });

    const document = render(<RecordingsView />);
    expect(document.asFragment()).toMatchSnapshot();
    expect(useFetchRecordings).toHaveBeenCalledWith({
      status: expectedStatus,
      sort_key: 'startsAt',
      limit: 50,
      offset: 0,
    });
  }
);

test.each([
  [null, 'upcoming'],
  ['upcoming', 'upcoming'],
  ['finished', 'finished'],
  ['failed', 'failed'],
  ['removed', 'removed'],
])(
  'should render loading: status=%s',
  (status: string | null, expectedStatus: string) => {
    mockStatus(status);
    vi.mocked(useLoading).mockReturnValue({
      isLoading: true,
      setIsLoading: vi.fn(),
    });

    const document = render(<RecordingsView />);
    expect(document.asFragment()).toMatchSnapshot();
    expect(useFetchRecordings).toHaveBeenCalledWith({
      status: expectedStatus,
      sort_key: 'startsAt',
      limit: 50,
      offset: 0,
    });
  }
);

test.each([
  [null, 'upcoming'],
  ['upcoming', 'upcoming'],
  ['finished', 'finished'],
  ['failed', 'failed'],
  ['removed', 'removed'],
])(
  'should render with no recordings: status=%s',
  (status: string | null, expectedStatus: string) => {
    mockStatus(status);
    const document = render(<RecordingsView />);
    expect(document.asFragment()).toMatchSnapshot();
    expect(useFetchRecordings).toHaveBeenCalledWith({
      status: expectedStatus,
      sort_key: 'startsAt',
      limit: 50,
      offset: 0,
    });
  }
);

test.each([
  [null, 'upcoming'],
  ['upcoming', 'upcoming'],
  ['finished', 'finished'],
  ['failed', 'failed'],
  ['removed', 'removed'],
])(
  'should render single page: status=%s',
  (status: string | null, expectedStatus: string) => {
    mockStatus(status);
    const recordings = buildRecordings(expectedStatus, 5);

    vi.mocked(useFetchRecordings).mockReturnValue({
      recordings,
      error: null,
      fetch: vi.fn(),
      total: 10,
    });

    const document = render(<RecordingsView />);
    expect(document.asFragment()).toMatchSnapshot();
    expect(useFetchRecordings).toHaveBeenCalledWith({
      status: expectedStatus,
      sort_key: 'startsAt',
      limit: 50,
      offset: 0,
    });
  }
);

test.each([
  [null, 'upcoming'],
  ['upcoming', 'upcoming'],
  ['finished', 'finished'],
  ['failed', 'failed'],
  ['removed', 'removed'],
])(
  'should render multiple pages: status=%s',
  (status: string | null, expectedStatus: string) => {
    mockStatus(status);
    const recordings = buildRecordings(expectedStatus, 5);

    vi.mocked(useFetchRecordings).mockReturnValue({
      recordings,
      error: null,
      fetch: vi.fn(),
      total: 150,
    });

    const document = render(<RecordingsView />);
    expect(document.asFragment()).toMatchSnapshot();
    expect(useFetchRecordings).toHaveBeenCalledWith({
      status: expectedStatus,
      sort_key: 'startsAt',
      limit: 50,
      offset: 0,
    });
  }
);

describe('cancel recordings', () => {
  test('should cancel all recordings', async () => {
    const recordings = buildRecordings('upcoming', 5);
    // Set first recording to recording state.
    recordings[0].status = 'recording';

    vi.mocked(useFetchRecordings).mockReturnValue({
      recordings,
      error: null,
      fetch: fetchMock,
      total: 150,
    });

    const document = render(<RecordingsView />);

    // Get elements by test id.
    const selectAllCheckbox = document.getByTestId(
      TestIds.SELECT_ALL_RECORDINGS_CHECKBOX
    );
    const cancelButton = document.getByTestId(
      TestIds.DELETE_CANCEL_RECORDINGS_BUTTON
    );
    const confirmButton = document.getByTestId(TestIds.CONFIRM_DELETE_BUTTON);

    await userEvent.click(selectAllCheckbox);
    expect(selectAllCheckbox).toBeChecked();

    const checkboxes = document.getAllByTestId(
      TestIds.SELECT_RECORDING_CHECKBOX
    );
    expect(checkboxes).toHaveLength(5);
    checkboxes.forEach((checkbox) => expect(checkbox).toBeChecked());

    await userEvent.click(cancelButton);
    await userEvent.click(confirmButton);

    expect(stopAndCancelRecordingsMock).toHaveBeenCalledWith(
      [recordings[0].id],
      recordings.slice(1).map((recording) => recording.id)
    );

    expect(selectAllCheckbox).not.toBeChecked();
    expect(fetchMock).toHaveBeenCalled();
  });

  test('should cancel selected recordings', async () => {
    const recordings = buildRecordings('upcoming', 5);
    // Set first recording to recording state.
    recordings[0].status = 'recording';

    vi.mocked(useFetchRecordings).mockReturnValue({
      recordings,
      error: null,
      fetch: fetchMock,
      total: 150,
    });

    const document = render(<RecordingsView />);

    // Get elements by test id.
    const selectAllCheckbox = document.getByTestId(
      TestIds.SELECT_ALL_RECORDINGS_CHECKBOX
    );
    const recordingsCheckboxes = document.getAllByTestId(
      TestIds.SELECT_RECORDING_CHECKBOX
    );
    const cancelButton = document.getByTestId(
      TestIds.DELETE_CANCEL_RECORDINGS_BUTTON
    );
    const confirmButton = document.getByTestId(TestIds.CONFIRM_DELETE_BUTTON);

    await userEvent.click(recordingsCheckboxes[0]);
    expect(recordingsCheckboxes[0]).toBeChecked();

    await userEvent.click(recordingsCheckboxes[1]);
    expect(recordingsCheckboxes[1]).toBeChecked();

    await userEvent.click(cancelButton);
    await userEvent.click(confirmButton);

    expect(selectAllCheckbox).not.toBeChecked();
    expect(stopAndCancelRecordingsMock).toHaveBeenCalledWith(
      [recordings[0].id],
      [recordings[1].id]
    );
  });
});

describe('delete recordings', () => {
  test('should delete all recordings', async () => {
    mockStatus('finished');
    const recordings = buildRecordings('finished', 5);

    vi.mocked(useFetchRecordings).mockReturnValue({
      recordings,
      error: null,
      fetch: fetchMock,
      total: 150,
    });

    const document = render(<RecordingsView />);

    // Get elements by test id.
    const selectAllCheckbox = document.getByTestId(
      TestIds.SELECT_ALL_RECORDINGS_CHECKBOX
    );
    const deleteButton = document.getByTestId(
      TestIds.DELETE_CANCEL_RECORDINGS_BUTTON
    );
    const confirmButton = document.getByTestId(TestIds.CONFIRM_DELETE_BUTTON);

    await userEvent.click(selectAllCheckbox);
    expect(selectAllCheckbox).toBeChecked();

    const checkboxes = document.getAllByTestId(
      TestIds.SELECT_RECORDING_CHECKBOX
    );
    expect(checkboxes).toHaveLength(5);
    checkboxes.forEach((checkbox) => expect(checkbox).toBeChecked());

    await userEvent.click(deleteButton);
    await userEvent.click(confirmButton);

    expect(removeRecordingsMock).toHaveBeenCalledWith(
      recordings.map((recording) => recording.id)
    );

    expect(selectAllCheckbox).not.toBeChecked();
    expect(fetchMock).toHaveBeenCalled();
  });

  test('should delete selected recordings', async () => {
    mockStatus('finished');
    const recordings = buildRecordings('finished', 5);

    vi.mocked(useFetchRecordings).mockReturnValue({
      recordings,
      error: null,
      fetch: fetchMock,
      total: 150,
    });

    const document = render(<RecordingsView />);

    // Get elements by test id.
    const selectAllCheckbox = document.getByTestId(
      TestIds.SELECT_ALL_RECORDINGS_CHECKBOX
    );
    const recordingsCheckboxes = document.getAllByTestId(
      TestIds.SELECT_RECORDING_CHECKBOX
    );
    const deleteButton = document.getByTestId(
      TestIds.DELETE_CANCEL_RECORDINGS_BUTTON
    );
    const confirmButton = document.getByTestId(TestIds.CONFIRM_DELETE_BUTTON);

    await userEvent.click(recordingsCheckboxes[0]);
    expect(recordingsCheckboxes[0]).toBeChecked();

    await userEvent.click(recordingsCheckboxes[1]);
    expect(recordingsCheckboxes[1]).toBeChecked();

    await userEvent.click(deleteButton);
    await userEvent.click(confirmButton);

    expect(selectAllCheckbox).not.toBeChecked();

    expect(removeRecordingsMock).toHaveBeenCalledWith([
      recordings[0].id,
      recordings[1].id,
    ]);
  });
});

describe('change status', () => {
  test.each([
    ['upcoming', 'finished'],
    ['upcoming', 'failed'],
    ['upcoming', 'removed'],
    ['finished', 'upcoming'],
    ['finished', 'failed'],
    ['finished', 'removed'],
    ['failed', 'upcoming'],
    ['failed', 'finished'],
    ['failed', 'removed'],
    ['removed', 'upcoming'],
    ['removed', 'finished'],
    ['removed', 'failed'],
  ])(
    "should change status from '%s' to'%s'",
    async (currentStatus: string, newStatus: string) => {
      mockStatus(currentStatus);
      const recordings = buildRecordings(currentStatus, 5);

      vi.mocked(useFetchRecordings).mockReturnValue({
        recordings,
        error: null,
        fetch: fetchMock,
        total: 150,
      });

      const document = render(<RecordingsView />);

      // Get elements by test id.
      const selectAllCheckbox = document.getByTestId(
        TestIds.SELECT_ALL_RECORDINGS_CHECKBOX
      );

      const statusDropdown = document.getByTestId(
        TestIds.RECORDINGS_STATUS_DROPDOWN
      );

      await userEvent.click(selectAllCheckbox);
      expect(selectAllCheckbox).toBeChecked();

      await userEvent.selectOptions(statusDropdown, newStatus);

      expect(setQueryParamsMock).toHaveBeenCalledWith({
        status: newStatus,
      });

      // Verify that all checkboxes are unchecked after changing status.
      expect(selectAllCheckbox).not.toBeChecked();
    }
  );
});

const mockStatus = (status: string | null) => {
  const params = new URLSearchParams();
  status && params.set('status', status);
  vi.mocked(useSearchParams).mockReturnValue([params, setQueryParamsMock]);
};

const buildRecordings = (status: string, count: number): Recording[] => {
  return [...Array(count).keys()].map(
    (i): Recording => ({
      id: `${i + 1}`,
      channelId: 'someChannelId',
      channelName: 'Some Channel',
      createdAt: 0,
      description: 'Some Description',
      duration: 0,
      enabled: true,
      endPadding: 2,
      endsAt: 0,
      eventId: 1,
      extraText: 'Extra Text',
      filename: '',
      langTitle: {
        ger: 'Some Title',
      },
      originalEndsAt: 0,
      originalStartsAt: 0,
      piconId: 1,
      startPadding: 3,
      startsAt: 0,
      status,
      subtitle: 'Some Subtitle',
      title: 'Some Title',
    })
  );
};
