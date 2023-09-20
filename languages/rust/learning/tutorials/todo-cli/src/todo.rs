use std::collections::HashMap;

#[derive(Debug)]
pub(crate) struct Todo {
    task_list: HashMap<String, TodoItem>,
}

impl Todo {
    pub(crate) fn new() -> Self {
        Todo {
            task_list: HashMap::new(),
        }
    }
    pub(crate) fn insert(&mut self, key: String, note: Option<String>) {
        // adds a new TodoItem to our list, with an optional note
        self.task_list.insert(
            key,
            TodoItem {
                note,
                completed: false,
            },
        );
    }

    pub(crate) fn list_all(&self) {
        if self.task_list.is_empty() {
            println!("List is empty!");
        } else {
            self.task_list.iter().for_each(|a| {
                println!("{:40} - {:40} - {}", a.0, a.1.get_note(), a.1.get_status());
            });
        }
    }
}

#[derive(Debug)]
pub(crate) struct TodoItem {
    note: Option<String>,
    completed: bool,
}

impl TodoItem {
    fn get_note(&self) -> String {
        String::from(match &self.note {
            Some(msg) => msg,
            _ => "No note",
        })
    }

    fn set_completed(&mut self) {
        self.completed = true;
    }

    fn get_status(&self) -> String {
        String::from(match self.completed {
            true => "Done",
            false => "Pending"
        })
    }
}
