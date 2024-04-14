import { afterEach, beforeEach, describe, expect, test, vi } from 'vitest';
import { cleanup, render } from '@testing-library/react';
import { useLoaderData, useNavigate } from 'react-router-dom';

import { Recording } from '@/clients/api/api.types';
import { Component as RecordingDetailView } from './RecordingDetailView';
import { useManageRecordingByEvent } from '@/hooks/recording';
import userEvent from '@testing-library/user-event';

vi.mock('@/hooks/recording');
vi.mock('react-router-dom');

const stopRecordingMock = vi.fn();
const cancelRecordingMock = vi.fn();
const removeRecordingMock = vi.fn();
const updateRecordingMock = vi.fn();

const RECORDING_ID = 'someID';

const navigateMock = vi.fn();

beforeEach(() => {
  vi.mocked(useManageRecordingByEvent).mockReturnValue({
    pending: false,
    cancelRecording: cancelRecordingMock,
    createRecording: async () => {},
    removeRecording: removeRecordingMock,
    stopRecording: stopRecordingMock,
    updateRecording: updateRecordingMock,
  });

  vi.mocked(useNavigate).mockReturnValue(navigateMock);
});

afterEach(() => {
  vi.resetAllMocks();
  cleanup();
});

describe('with status recording', () => {
  test('should render recording', async () => {
    const recording = buildRecording('recording');

    vi.mocked(useLoaderData).mockReturnValue(recording);

    const document = render(<RecordingDetailView />);

    expect(document.asFragment()).toMatchSnapshot();
  });

  test('should stop recording', async () => {
    const recording = buildRecording('recording');

    vi.mocked(useLoaderData).mockReturnValue(recording);

    const document = render(<RecordingDetailView />);

    await userEvent.click(document.getByText('stop_recording'));
    await document.findByText('confirm_stop_recording');

    expect(document.getByText('confirm_stop_recording')).toBeInTheDocument();

    await userEvent.click(document.getByText('stop'));

    expect(stopRecordingMock).toHaveBeenCalledOnce();
  });

  test('should update a recording', async () => {
    const recording = buildRecording('recording');
    vi.mocked(useLoaderData).mockReturnValue(recording);

    const document = render(<RecordingDetailView />);

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
    vi.mocked(useLoaderData).mockReturnValue(recording);

    const document = render(<RecordingDetailView />);

    expect(document.asFragment()).toMatchSnapshot();
  });

  test('should delete recording', async () => {
    const recording = buildRecording('completed');

    vi.mocked(useLoaderData).mockReturnValue(recording);

    const document = render(<RecordingDetailView />);

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

    vi.mocked(useLoaderData).mockReturnValue(recording);

    const document = render(<RecordingDetailView />);

    expect(document.asFragment()).toMatchSnapshot();
  });

  test('should cancel recording', async () => {
    const recording = buildRecording('scheduled');

    vi.mocked(useLoaderData).mockReturnValue(recording);

    const document = render(<RecordingDetailView />);

    await userEvent.click(document.getByText('cancel_recording'));
    await document.findByText('confirm_cancel_recording');

    expect(document.getByText('confirm_cancel_recording')).toBeInTheDocument();

    await userEvent.click(document.getByText('cancel'));

    expect(cancelRecordingMock).toHaveBeenCalledOnce();
  });

  test('should update a recording', async () => {
    const recording = buildRecording('scheduled');

    vi.mocked(useLoaderData).mockReturnValue(recording);

    const document = render(<RecordingDetailView />);

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
