import DashboardLayout from '../layout/DashboardLayout.vue'
// GeneralViews
import NotFound from '../pages/NotFoundPage.vue'

// Admin pages
import Courses from 'src/pages/Courses.vue'
import UserProfile from 'src/pages/UserProfile.vue'
import TableList from 'src/pages/TableList.vue'
import Typography from 'src/pages/Typography.vue'
import Icons from 'src/pages/Icons.vue'
import Maps from 'src/pages/Maps.vue'
import Notifications from 'src/pages/Notifications.vue'
import Upgrade from 'src/pages/Upgrade.vue'
import Course from 'src/pages/Course.vue'
import CourseVideos from 'src/pages/CourseVideos.vue'
import Quiz from 'src/components/Quiz.vue'
import CourseListOfVideos from 'src/pages/CourseListOfVideos.vue'
import BeforePayment from 'src/pages/BeforePayment.vue'
import LogIn from 'src/pages/LogIn.vue'
import Restore from 'src/pages/Restore.vue'
import EmailSent from 'src/pages/EmailSent.vue'
import RecovPassLoggedIn from 'src/pages/RecovPassLoggedIn.vue'
import AddCourse from 'src/pages/AddCourse.vue'
import ListOfCourses from 'src/pages/Admin/ListOfCourses.vue'
import ListOfCustomers from 'src/pages/Admin/ListOfCustomers.vue'
import ListOfOrders from 'src/pages/Admin/ListOfOrders'

const routes = [
  {
    path: '/',
    component: DashboardLayout,
    redirect: '/admin/courses',
    meta: {
      requiresAuth: true
    }
  },

  {
    path: '/admin',
    component: DashboardLayout,
    redirect: '/admin/courses',
    children: [
      {
        path: 'courses',
        name: 'Courses',
        component: Courses,
        meta: {
          requiresAuth: true
        }
      },
      {
        path: 'user',
        name: 'User',
        component: UserProfile,
        meta: {
          requiresAuth: true
        }
      },

      {
        path: 'courses/course/:id_course',
        name: 'Course',
        component: Course,
        meta: {
          requiresAuth: true
        }
      },
      {
        path: 'courses/course/:id_course/list_of_videos/videos/',
        name: 'CourseVideos',
        component: CourseVideos,
        meta: {
          requiresAuth: true
        }
      },
      {
        path: 'courses/course/:id_course/list_of_videos',
        name: 'CourseListOfVideos',
        component: CourseListOfVideos,
        meta: {
          requiresAuth: true
        }
      },
      {
        path: 'courses/course/before_payment',
        name: 'BeforePayment',
        component: BeforePayment,
        meta: {
          requiresAuth: true
        }
      },
      {
        path: '/log-in',
        name: 'LogIn',
        component: LogIn
      },
      {
        path: '/restore',
        name: 'Restore',
        component: Restore
      },
      {
        path: '/email_sent',
        name: 'EmailSent',
        component: EmailSent
      },
      {
        path: '/recover_pass',
        name: 'RecovPassLoggedIn',
        component: RecovPassLoggedIn,
        meta: {
          requiresAuth: true
        }
      },
      // ADMIN PANEL

      {
        path: '/add_course',
        name: 'AddCourse',
        component: AddCourse,
        meta: {
          requiresAuth: true
        }
      },

      {
        path: '/list_of_courses',
        name: 'ListOfCourses',
        component: ListOfCourses,
        meta: {
          requiresAuth: true
        }
      },
      {
        path: '/list_of_customers',
        name: 'ListOfCustomers',
        component: ListOfCustomers,
        meta: {
          requiresAuth: true
        }
      },
      {
        path: '/list_of_orders',
        name: 'ListOfOrders',
        component: ListOfOrders,
        meta: {
          requiresAuth: true
        }
      },
    ]
  },

  { path: '*', component: NotFound }

]


/**
 * Asynchronously load view (Webpack Lazy loading compatible)
 * The specified component must be inside the Views folder
 * @param  {string} name  the filename (basename) of the view to load.
function view(name) {
   var res= require('../components/Dashboard/Views/' + name + '.vue');
   return res;
};**/

export default routes
