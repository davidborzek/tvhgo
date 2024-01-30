import { render, screen, cleanup } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { afterEach, beforeEach, describe, expect, test, vi } from 'vitest';
import RecordingDetailView from './RecordingDetailView';
import { MemoryRouter, Route, Routes } from 'react-router-dom';
import {
  useFetchRecording,
  useManageRecordingByEvent,
} from '@/hooks/recording';
import { Recording } from '@/clients/api/api.types';
import { PropsWithChildren } from 'react';

vi.mock('@/hooks/recording');

const stopRecordingMock = vi.fn();
const cancelRecordingMock = vi.fn();
const removeRecordingMock = vi.fn();
const updateRecordingMock = vi.fn();

const RECORDING_ID = 'someID';

beforeEach(() => {
  vi.mocked(useManageRecordingByEvent).mockReturnValue({
    pending: false,
    cancelRecording: cancelRecordingMock,
    createRecording: async () => {},
    removeRecording: removeRecordingMock,
    stopRecording: stopRecordingMock,
    updateRecording: updateRecordingMock,
  });
});

afterEach(() => {
  vi.resetAllMocks();
  cleanup();
});

test('should render error', () => {
  const fetchMock = vi.fn();

  const error = 'some error occurred';
  vi.mocked(useFetchRecording).mockReturnValue({
    error,
    recording: undefined,
    fetch: fetchMock,
  });

  render(<RecordingDetailView />, { wrapper: TestRouter });

  const errorDiv = screen.getByText(error);
  expect(errorDiv).toBeInTheDocument();

  expect(fetchMock).toHaveBeenNthCalledWith(1, RECORDING_ID);
});

describe('with status recording', () => {
  test('should render recording', async () => {
    const recording = buildRecording('recording');

    const fetchMock = vi.fn();
    vi.mocked(useFetchRecording).mockReturnValue({
      error: null,
      recording,
      fetch: fetchMock,
    });

    const document = render(<RecordingDetailView />, { wrapper: TestRouter });

    expect(document.asFragment()).toMatchSnapshot();
    expect(fetchMock).toHaveBeenNthCalledWith(1, RECORDING_ID);
  });

  test('should stop recording', async () => {
    const recording = buildRecording('recording');

    vi.mocked(useFetchRecording).mockReturnValue({
      error: null,
      recording,
      fetch: async () => {},
    });

    const document = render(<RecordingDetailView />, { wrapper: TestRouter });

    await userEvent.click(document.getByText('stop_recording'));
    await document.findByText('confirm_stop_recording');

    expect(document.getByText('confirm_stop_recording')).toBeInTheDocument();

    await userEvent.click(document.getByText('stop'));

    expect(stopRecordingMock).toHaveBeenCalledOnce();
  });

  test('should update a recording', async () => {
    const recording = buildRecording('recording');

    vi.mocked(useFetchRecording).mockReturnValue({
      error: null,
      recording,
      fetch: async () => {},
    });

    const document = render(<RecordingDetailView />, { wrapper: TestRouter });

    const endPaddingInput = document.container.querySelector(
      'input[name="endPadding"]'
    ) as Element;

    const endPadding = 234;

    await userEvent.type(endPaddingInput, `${endPadding}`);

    expect(endPaddingInput).toHaveValue(endPadding);

    await userEvent.click(document.getByText('save'));

    expect(updateRecordingMock).toHaveBeenNthCalledWith(1, recording.id, {
      startPadding: recording.startPadding,
      endPadding,
    });
  });
});

describe('with status completed', () => {
  test('should render recording', async () => {
    const recording = buildRecording('completed');

    const fetchMock = vi.fn();
    vi.mocked(useFetchRecording).mockReturnValue({
      error: null,
      recording,
      fetch: fetchMock,
    });

    const document = render(<RecordingDetailView />, { wrapper: TestRouter });

    expect(document.asFragment()).toMatchSnapshot();
    expect(fetchMock).toHaveBeenNthCalledWith(1, RECORDING_ID);
  });

  test('should delete recording', async () => {
    const recording = buildRecording('completed');

    vi.mocked(useFetchRecording).mockReturnValue({
      error: null,
      recording,
      fetch: async () => {},
    });

    const document = render(<RecordingDetailView />, { wrapper: TestRouter });

    await userEvent.click(document.getByText('delete_recording'));
    await document.findByText('confirm_delete_recording');

    expect(document.getByText('confirm_delete_recording')).toBeInTheDocument();

    await userEvent.click(document.getByText('delete'));

    expect(removeRecordingMock).toHaveBeenCalledOnce();
  });
});

describe('with status scheduled', () => {
  test('should render recording', async () => {
    const recording = buildRecording('scheduled');

    const fetchMock = vi.fn();
    vi.mocked(useFetchRecording).mockReturnValue({
      error: null,
      recording,
      fetch: fetchMock,
    });

    const document = render(<RecordingDetailView />, { wrapper: TestRouter });

    expect(document.asFragment()).toMatchSnapshot();
    expect(fetchMock).toHaveBeenNthCalledWith(1, RECORDING_ID);
  });

  test('should cancel recording', async () => {
    const recording = buildRecording('scheduled');

    vi.mocked(useFetchRecording).mockReturnValue({
      error: null,
      recording,
      fetch: async () => {},
    });

    const document = render(<RecordingDetailView />, { wrapper: TestRouter });

    await userEvent.click(document.getByText('cancel_recording'));
    await document.findByText('confirm_cancel_recording');

    expect(document.getByText('confirm_cancel_recording')).toBeInTheDocument();

    await userEvent.click(document.getByText('cancel'));

    expect(cancelRecordingMock).toHaveBeenCalledOnce();
  });

  test('should update a recording', async () => {
    const recording = buildRecording('scheduled');

    vi.mocked(useFetchRecording).mockReturnValue({
      error: null,
      recording,
      fetch: async () => {},
    });

    const document = render(<RecordingDetailView />, { wrapper: TestRouter });

    const startPaddingInput = document.container.querySelector(
      'input[name="startPadding"]'
    ) as Element;
    const endPaddingInput = document.container.querySelector(
      'input[name="endPadding"]'
    ) as Element;

    const startPadding = 123;
    const endPadding = 234;

    await userEvent.type(startPaddingInput, `${startPadding}`);
    await userEvent.type(endPaddingInput, `${endPadding}`);

    expect(startPaddingInput).toHaveValue(startPadding);
    expect(endPaddingInput).toHaveValue(endPadding);

    await userEvent.click(document.getByText('save'));

    expect(updateRecordingMock).toHaveBeenNthCalledWith(1, recording.id, {
      startPadding,
      endPadding,
    });
  });
});

const buildRecording = (status: string): Recording => {
  return {
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
    id: RECORDING_ID,
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
  };
};

const TestRouter = ({ children }: PropsWithChildren) => {
  return (
    <MemoryRouter initialEntries={[`/recordings/${RECORDING_ID}`]}>
      <Routes>
        <Route path="/recordings/:id" element={children} />
      </Routes>
    </MemoryRouter>
  );
};
