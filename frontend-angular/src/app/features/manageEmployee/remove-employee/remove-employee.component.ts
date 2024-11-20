import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA, MatDialogRef } from '@angular/material/dialog';
import { EmployeeService } from '../../../services/employee.service';
import { Employee } from '../../../models/employee.model';

@Component({
  selector: 'app-remove-employee',
  templateUrl: './remove-employee.component.html',
  styleUrl: './remove-employee.component.css'
})
export class RemoveEmployeeComponent {

  constructor(
      public dialogRef: MatDialogRef<RemoveEmployeeComponent>,
      private employeeService: EmployeeService,
      @Inject(MAT_DIALOG_DATA) public data: { employee: Employee } 
  ) {}

    removeEmployee():void{
      this.employeeService.delete(this.data.employee.id).subscribe({
        next: (data) => {
          console.log('Employee Removed successfully:', data);
          this.dialogRef.close(data);  
        },
        error: (e) => {
          console.error('Error Removing employee:', e);
        }
      });
  }
  
  onCancel(): void {
    this.dialogRef.close();
  }

  close(): void {
    this.dialogRef.close();
  }

}
