import { FormikValues, useFormik } from "formik";
import { RefObject, useEffect } from "react";

const useFormikErrorFocus = <T extends FormikValues>(
  formik: ReturnType<typeof useFormik<T>>,
  ...refs: RefObject<HTMLInputElement>[]
) => {
  useEffect(() => {
    if (formik.isSubmitting) {
      for (const [field, error] of Object.entries(formik.errors)) {
        const ref = refs.find((ref) => ref.current?.name == field);

        if (ref && error) {
          ref.current?.focus();
          return;
        }
      }
    }
  }, [formik.isSubmitting]);
};

export default useFormikErrorFocus;
