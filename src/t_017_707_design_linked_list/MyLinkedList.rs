struct MyLinkedList {}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyLinkedList {
    fn new() -> Self {
        MyLinkedList {}
    }

    fn get(&self, index: i32) -> i32 {
        1
    }

    fn add_at_head(&self, val: i32) {}

    fn add_at_tail(&self, val: i32) {}

    fn add_at_index(&self, index: i32, val: i32) {}

    fn delete_at_index(&self, index: i32) {}
}
