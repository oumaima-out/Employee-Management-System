
import { AfterViewInit, Component, OnInit, ViewChild } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatPaginator } from '@angular/material/paginator';
import { MatSort } from '@angular/material/sort';
import { Employee } from '../../../models/employee.model';
import { EmployeeService } from '../../../services/employee.service';
import { MatDialog } from '@angular/material/dialog';
import { AddEmployeeComponent } from '../add-employee/add-employee.component';
import { SelectionModel } from '@angular/cdk/collections';
import { EditEmployeeComponent } from '../edit-employee/edit-employee.component';
import { RemoveEmployeeComponent } from '../remove-employee/remove-employee.component';

@Component({
  selector: 'app-employees',
  templateUrl: './employees.component.html',
  styleUrls: ['./employees.component.css']
})

export class EmployeesComponent implements OnInit, AfterViewInit {
  
  public employees?: Employee[];

  public displayedColumns = ['select' ,'id', 'first_name', 'last_name', 'email', 'phone', 'department', 'date_of_hire', 'position']; 
  public dataSource: MatTableDataSource<any>;
  public selection = new SelectionModel<Employee>(true, []);

  @ViewChild(MatPaginator) paginator!: MatPaginator;
  @ViewChild(MatSort) sort!: MatSort;

  constructor(private employeeService: EmployeeService, private dialog: MatDialog) {
    this.dataSource = new MatTableDataSource(this.employees);
  }

  ngOnInit() {
    this.retrieveEmployees();
  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  retrieveEmployees(): void {
    this.employeeService.getAll().subscribe({
      next: (data) => {
        this.employees = data;
        this.dataSource = new MatTableDataSource(this.employees);
      },
      error: (e) => console.error(e)
    });
  }
  
  addEmployeeDialog(): void {
    const dialogRef = this.dialog.open(AddEmployeeComponent, {
      width: '600px', 
      height:'500px'
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log('Le modal a été fermé avec la donnée:', result);
      this.retrieveEmployees();
    });
  }

  editEmployeeDialog(): void {
    if (this.selection.selected.length === 0 || this.selection.selected.length>2) {
      alert('Please select one employee to edit.');
      return;
    }

    const selectedEmployee = this.selection.selected[0];

    const dialogRef = this.dialog.open(EditEmployeeComponent, {
      width: '600px', 
      height:'500px',
      data: { employee: selectedEmployee }
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log('Le modal a été fermé avec la donnée:', result);
      this.retrieveEmployees();
    });
  }
  
  removeEmployeeDialog():void {
    if (this.selection.selected.length === 0 || this.selection.selected.length>2) {
      alert('Please select one employee to remove.');
      return;
    }

    const selectedEmployee = this.selection.selected[0];

    const dialogRef = this.dialog.open(RemoveEmployeeComponent, {
      width: '400px', 
      height:'200px',
      data: { employee: selectedEmployee }
    });

    dialogRef.afterClosed().subscribe(result => {
      console.log('Le modal a été fermé avec la donnée:', result);
      this.retrieveEmployees();
    });
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.dataSource.data.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  toggleAllRows() {
    if (this.isAllSelected()) {
      this.selection.clear();
      return;
    }

    this.selection.select(...this.dataSource.data);
  }

  /** The label for the checkbox on the passed row */
  checkboxLabel(row?: Employee): string {
    if (!row) {
      return `${this.isAllSelected() ? 'deselect' : 'select'} all`;
    }
    return `${this.selection.isSelected(row) ? 'deselect' : 'select'} row ${row.id}`;  
  }

}
