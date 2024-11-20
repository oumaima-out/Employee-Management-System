import { Component, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { EmployeeService } from '../../../services/employee.service';

@Component({
  selector: 'app-add-employee',
  templateUrl: './add-employee.component.html',
  styleUrls: ['./add-employee.component.css']
})

export class AddEmployeeComponent {

  employeeForm: FormGroup;

  constructor(
    private fb: FormBuilder,
    public dialogRef: MatDialogRef<AddEmployeeComponent>,
    private employeeService: EmployeeService
  ) {
    this.employeeForm = this.fb.group({
      first_name: ['', Validators.required],
      last_name: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      phone: ['', Validators.required],
      department: ['', Validators.required],
      date_of_hire: ['', Validators.required],
      position: ['', Validators.required]
    });
  }

  onSubmit(): void {
    if (this.employeeForm.valid) {
      this.employeeService.create(this.employeeForm.value).subscribe({
        next: (data) => {
          console.log('Employee created successfully:', data);
          this.dialogRef.close(data);  
        },
        error: (e) => {
          console.error('Error creating employee:', e);
        }
      });
    } else {
      console.log('Form is invalid');
    }
  }
  
  onCancel(): void {
    this.dialogRef.close();
  }

  close(): void {
    this.dialogRef.close();
  }
}
