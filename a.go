//原因：lock是在子线程里，而unlock是在主线程里，并且主线程没有被lock