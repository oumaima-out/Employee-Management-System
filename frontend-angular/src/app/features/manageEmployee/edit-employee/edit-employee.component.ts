import { Component, Inject } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { EmployeeService } from '../../../services/employee.service';
import { Employee } from '../../../models/employee.model';

@Component({
  selector: 'app-edit-employee',
  templateUrl: './edit-employee.component.html',
  styleUrl: './edit-employee.component.css'
})

export class EditEmployeeComponent {

  employeeForm: FormGroup;

  constructor(
    private fb: FormBuilder,
    public dialogRef: MatDialogRef<EditEmployeeComponent>,
    private employeeService: EmployeeService,
    @Inject(MAT_DIALOG_DATA) public data: { employee: Employee } 
  ) {
    this.employeeForm = this.fb.group({
      id: [data.employee.id],
      first_name: [data.employee.first_name, Validators.required],
      last_name: [data.employee.last_name, Validators.required],
      email: [data.employee.email, [Validators.required, Validators.email]],
      phone: [data.employee.phone, Validators.required],
      department: [data.employee.department, Validators.required],
      date_of_hire: [data.employee.date_of_hire, Validators.required],
      position: [data.employee.position, Validators.required],
    });
  }

  onSubmit(): void {
    if (this.employeeForm.valid) {
      this.employeeService.update(this.data.employee.id,this.employeeForm.value).subscribe({
        next: (data) => {
          console.log('Employee Edited successfully:', data);
          this.dialogRef.close(data);  
        },
        error: (e) => {
          console.error('Error updating employee:', e);
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
