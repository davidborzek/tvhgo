import { FormikValues, useFormik } from 'formik';
import { RefObject, useEffect } from 'react';

const useFormikErrorFocus = <T extends FormikValues>(
  { isSubmitting, errors }: ReturnType<typeof useFormik<T>>,
  ...refs: RefObject<HTMLInputElement>[]
) => {
  useEffect(() => {
    if (isSubmitting) {
      for (const [field, error] of Object.entries(errors)) {
        const ref = refs.find((ref) => ref.current?.name === field);

        if (ref && error) {
          ref.current?.focus();
          return;
        }
      }
    }
  }, [isSubmitting, errors, refs]);
};

export default useFormikErrorFocus;
